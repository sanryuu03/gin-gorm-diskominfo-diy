ini soal nomor satu
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

ini soal nomor 2
SELECT
    r.id,
    r.name,
    r.capacity,
    r.is_active
FROM rooms r
WHERE r.is_active = true
  AND NOT EXISTS (
      SELECT 1
      FROM bookings b
      WHERE CAST(b.room_id AS BIGINT) = r.id
        AND b.status = 'CONFIRMED'
        AND b.start_time < TIMESTAMP '2026-12-01 12:00:00'
        AND b.end_time   > TIMESTAMP '2026-12-01 10:00:00'
  )
ORDER BY r.id;