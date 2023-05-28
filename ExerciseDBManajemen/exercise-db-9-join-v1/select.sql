
SELECT
reports.id AS id,
students.fullname AS fullname,
students.class AS class,
students.status AS status,
reports.study AS study,
reports.score AS score
FROM reports
INNER JOIN students
ON reports.student_id = students.id
WHERE reports.score < 70 AND students.status = 'active'
ORDER BY reports.score ASC;