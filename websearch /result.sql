CREATE TABLE IF NOT EXISTS search_results (
    url TEXT PRIMARY KEY,
    search_json TEXT NOT NULL,
    summary TEXT
);