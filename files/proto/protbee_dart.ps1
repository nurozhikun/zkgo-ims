# C:/protobuf/bin/protoc.exe -I"./protbee/" --go_out="../../zk-atom/zipbf/protbee" "./protbee/prot_bee_base.proto" 
$name = "C:/protobuf/bin/protoc.exe "
$name += ' -I"./protbee/"'
$name += ' --dart_out="D:/works/zkflutter/zkfly_ims/lib/pbf/protbee"'
$name += ' "./protbee/prot_bee_base.proto"'
$name += ' "./protbee/prot_bee_body.proto"'
$name += ' "./protbee/prot_bee_appdart.proto"'
echo $name
cmd /C $name
#  --dart_out="../../../sieflutter/fzksie/lib/pbf" 
# --cpp_out="../../../../geminicpp/src/pbf/robhost"
# --dart_out="../../../sieflutter/fzksie/lib/pbf"