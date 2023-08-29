package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneOutgoingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M14.3756 9.79424C15.1022 9.79424 15.7716 10.1883 16.1242 10.8235L18.5707 15.2303C18.891 15.8073 18.906 16.5052 18.6109 17.0955L16.2541 21.8092C16.2541 21.8092 16.937 25.3206 19.7955 28.179C22.6539 31.0374 26.1536 31.7086 26.1536 31.7086L30.8665 29.3522C31.4572 29.0568 32.1556 29.0721 32.7328 29.393L37.1521 31.85C37.7867 32.2028 38.1803 32.8719 38.1803 33.598V38.6715C38.1803 41.2552 35.7804 43.1213 33.3323 42.2952C28.3044 40.5987 20.4997 37.3685 15.5528 32.4216C10.606 27.4748 7.37577 19.6701 5.67922 14.6422C4.8532 12.1941 6.71929 9.79424 9.30297 9.79424H14.3756Z"/><path stroke-linecap="round" d="M35 6L43 14L35 22"/><path stroke-linecap="round" d="M27 14.0083H43"/></g>`),
		g.Group(children),
	)
}