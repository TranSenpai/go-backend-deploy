
# So sánh folder schema và folder migration
make ent

atlas migrate diff migration_name \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "docker://mysql/8/ent"

# Tạo lại atlas.sum: khi muốn tự thay đổi file sql migration
atlas migrate hash \
  --dir "file://ent/migrate/migrations" 


# Áp dụng migration vào db
make ent

atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url "mysql://root:12345@localhost:3307/go-backend"

# DB đã có sẵn
    - tạo init
    - bỏ qua không chạy init

atlas migrate set <version> \
  --dir "file://ent/migrate/migrations" \
  --url "mysql://root:12345@localhost:3307/go-backend"
    
    