# Khoá chính:
    - Table User:
        edge.To("Articles", Articles.Type),

# Khoá phụ:
    - Table Article:
        edge.From("Users", Users.Type).Ref("Articles").Field("user_id").Unique().Required(),