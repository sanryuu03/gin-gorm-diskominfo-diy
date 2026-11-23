SELECT
    r.id AS room_id,
    r.name,
    COALESCE(
        SUM(EXTRACT(EPOCH FROM (b.end_time - b.start_time)) / 3600),
        0
    ) AS total_jam
FROM rooms r
LEFT JOIN bookings b
    ON CAST(b.room_id AS BIGINT) = r.id
   AND b.status = 'CONFIRMED'
   AND b.start_time >= date_trunc('month', CURRENT_DATE)
   AND b.start_time <  date_trunc('month', CURRENT_DATE) + INTERVAL '1 month'
GROUP BY r.id, r.name
ORDER BY r.id;