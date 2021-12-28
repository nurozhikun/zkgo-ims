# C:/protobuf/bin/protoc.exe -I"./protbee/" --go_out="../../zk-atom/zipbf/protbee" "./protbee/prot_bee_base.proto" 
$name = "C:/protobuf/bin/protoc.exe "
$name += ' -I"./protbee/"'
$name += ' --go_out="../../"'
$name += ' "./protbee/prot_bee_base.proto"'
$name += ' "./protbee/prot_bee_body_map.proto"'
echo $name
cmd /C $name
#  --dart_out="../../../sieflutter/fzksie/lib/pbf" 
# --cpp_out="../../../../geminicpp/src/pbf/robhost"
# --dart_out="../../../sieflutter/fzksie/lib/pbf"