# header(固定大小)(HeadMarker+Head)
| marker | cmd    | pkgtye | bodylen |
| ------ | ------ | ------ | ------- |
| 4bytes | 4bytes | 2bytes | 2bytes  |
# body(header中的bodylen)
自定扩展的格式
# tailer(固定大小)(TailMarker+Tail)
| check  | marker |
| ------ | ------ |
| 4bytes | 4bytes |