package fruit

import (
        "github.com/golang/protobuf/proto"

        pb "github.com/leeren/bazel-gazelle-help/fruit/fruit_go_proto" // If I use `pb "fruit/fruit.proto" gazelle fails
)

func CreateDate(month, day int32) (pb.Date, error) {
  var date *pb.Date
  date = &pb.Date{
    Month: proto.Int32(month),
    Day: proto.Int32(day),
  }
  return date, err
}
