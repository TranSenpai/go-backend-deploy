https://entgo.io/docs/interceptors/

# Bước 1: Bật feature
    vào file: go-backend/ent/generate.go
    thêm cờ --feature intercept,schema/snapshot
    //go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/execquery,intercept,schema/snapshot ./schema

# Bước 2: lấy code mẫu và dán vào file
    code mẫu: https://entgo.io/docs/interceptors/#soft-delete - tab Mixxin
    tạo file chứa code: go-backend/ent/soft-delete/sof_delete_mixin.go

# Bước 3: Nếu muốn đổi tên cột xoá mềm
- line 23: delete_time > deleted_at

- line 62: SetDeleteTime > SetDeletedAt

- line 70: SetDeleteTime > SetDeletedAt

# Bước 4: thêm import để fix lỗi circular import ở bước sau
    thêm vào file new ent: go-backend/internal/common/ent_client/ent_client.go
    `import _ "<project>/ent/runtime"`


# Bước 5: chạy lệnh `go generate ./ent` LÂN 1 để tạo package intercept


# Bước 6: mở file go-backend/ent/soft-delete/sof_delete_mixin.go 
    thêm các import còn thiếu:
        - intercept
        - gen "go-backend/ent"

# Bước 7: sử dụng mixin - bị lỗi `import cycle not allowed`
    lấy code mẫu: https://entgo.io/docs/interceptors/#soft-delete - tab Mixin Usage

# Bước 8: chạy lệnh `go generate ./ent` LÂN 2 để fix lỗi `import cycle not allowed`