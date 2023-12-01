

directory=$1

echo $directory/api/docs/docs.go;

$sed '/LeftDelim/d' $directory/api/docs/docs.go
$sed '/RightDelim/d' $directory/api/docs/docs.go
