ROOT_DIR=$(git rev-parse --show-toplevel)

SRC_DIR="$ROOT_DIR/prototype/data/protobuf/src"
DST_DIR="$ROOT_DIR/prototype/data/protobuf/dst"

mkdir -p "$DST_DIR/postpb"
protoc -I=$SRC_DIR --go_out="$DST_DIR/postpb" "$SRC_DIR/post.proto"

# source - https://github.com/favadi/protoc-go-inject-tag
protoc-go-inject-tag -input="$DST_DIR/postpb/post.pb.go"
