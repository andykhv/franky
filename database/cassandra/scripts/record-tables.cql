CREATE TABLE IF NOT EXISTS records.records_by_user_date (
    user_id text,
    played_date date,
    played_time time,
    playlist text,
    category text,
    song text,
    artist text,
    album text,
    duration smallint,
    PRIMARY KEY ((user_id, played_date), played_time, playlist, category, song, artist, album)
);

CREATE TABLE IF NOT EXISTS records.records_by_user_date_playlist (
    user_id text,
    played_date date,
    played_time time,
    playlist text,
    category text,
    song text,
    artist text,
    album text,
    duration smallint,
    PRIMARY KEY ((user_id, played_date, playlist), played_time, category, song, artist, album)
);

CREATE TABLE IF NOT EXISTS records.records_by_user_date_playlist_artist (
    user_id text,
    played_date date,
    played_time time,
    playlist text,
    category text,
    song text,
    artist text,
    album text,
    duration smallint,
    PRIMARY KEY ((user_id, played_date, playlist, artist), played_time, category, song, album)
);

CREATE TABLE IF NOT EXISTS records.records_by_user_date_playlist_artist_song (
    user_id text,
    played_date date,
    played_time time,
    playlist text,
    category text,
    song text,
    artist text,
    album text,
    duration smallint,
    PRIMARY KEY ((user_id, played_date, playlist, artist, song), played_time, category, album)
);

CREATE TABLE IF NOT EXISTS records.records_by_user_date_playlist_category (
    user_id text,
    played_date date,
    played_time time,
    playlist text,
    category text,
    song text,
    artist text,
    album text,
    duration smallint,
    PRIMARY KEY ((user_id, played_date, playlist, category), played_time, song, artist, album)
);

CREATE TABLE IF NOT EXISTS records.records_by_user_date_album (
    user_id text,
    played_date date,
    played_time time,
    playlist text,
    category text,
    song text,
    artist text,
    album text,
    duration smallint,
    PRIMARY KEY ((user_id, played_date, album), played_time, playlist, category, song, artist)
);

CREATE TABLE IF NOT EXISTS records.records_by_user_date_category (
    user_id text,
    played_date date,
    played_time time,
    playlist text,
    category text,
    song text,
    artist text,
    album text,
    duration smallint,
    PRIMARY KEY ((user_id, played_date, category), played_time, playlist, song, artist, album)
);
