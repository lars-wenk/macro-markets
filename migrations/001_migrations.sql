-- Users & Subs (simplified)
CREATE TABLE IF NOT EXISTS users (
id BIGSERIAL PRIMARY KEY,
email TEXT UNIQUE NOT NULL,
pass_hash TEXT NOT NULL,
role TEXT NOT NULL DEFAULT 'free',
created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


CREATE TABLE IF NOT EXISTS subscriptions (
id BIGSERIAL PRIMARY KEY,
user_id BIGINT NOT NULL REFERENCES users(id),
tier TEXT NOT NULL,
status TEXT NOT NULL,
renew_at TIMESTAMPTZ
);


-- Sources
CREATE TABLE IF NOT EXISTS sources (
id BIGSERIAL PRIMARY KEY,
type TEXT NOT NULL, -- 'x','rss','api'
handle TEXT,
url TEXT,
trust_score NUMERIC DEFAULT 0,
active BOOLEAN NOT NULL DEFAULT TRUE
);


-- Raw items (blob in S3 ideally; keep hash for dedupe)
CREATE TABLE IF NOT EXISTS raw_items (
id BIGSERIAL PRIMARY KEY,
source_id BIGINT REFERENCES sources(id),
fetched_at TIMESTAMPTZ NOT NULL DEFAULT now(),
payload_json JSONB NOT NULL,
content_hash TEXT UNIQUE
);


-- Articles (normalized)
CREATE TABLE IF NOT EXISTS articles (
id BIGSERIAL PRIMARY KEY,
title TEXT NOT NULL,
text TEXT NOT NULL,
source_id BIGINT REFERENCES sources(id),
published_at TIMESTAMPTZ,
url TEXT,
lang TEXT,
sentiment NUMERIC,
dedupe_group TEXT
);


-- Macro events (calendar)
CREATE TABLE IF NOT EXISTS macro_events (
id BIGSERIAL PRIMARY KEY,
region TEXT NOT NULL,
name TEXT NOT NULL,
ts_release TIMESTAMPTZ NOT NULL,
consensus NUMERIC,
previous NUMERIC,
actual NUMERIC,
surprise_score NUMERIC,
url TEXT
);


-- Indexes
CREATE INDEX IF NOT EXISTS idx_articles_published_at ON articles(published_at DESC);
CREATE INDEX IF NOT EXISTS idx_macro_events_ts ON macro_events(ts_release DESC);
CREATE INDEX IF NOT EXISTS idx_raw_items_hash ON raw_items(content_hash);