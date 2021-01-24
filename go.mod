module robotgo

go 1.14

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/image => github.com/golang/image v0.0.0-20200927104501-e162460cd6b5
	golang.org/x/net => github.com/golang/net v0.0.0-20201002202402-0a1ea396d57c
	golang.org/x/sys => github.com/golang/sys v0.0.0-20201018230417-eeed37f84f13
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/win => github.com/golang/win v0.0.0-20201021184640-9c48443b1f44
)

require (
	github.com/go-vgo/robotgo v0.92.1
	github.com/golang/protobuf v1.4.3
	golang.org/x/net v0.0.0-20201002202402-0a1ea396d57c
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)
