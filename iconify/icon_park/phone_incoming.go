package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneIncoming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M15.3758 9.79422C16.1023 9.79422 16.7717 10.1883 17.1244 10.8235L19.5708 15.2303C19.8911 15.8073 19.9062 16.5052 19.611 17.0955L17.2542 21.8092C17.2542 21.8092 17.9372 25.3206 20.7956 28.179C23.6541 31.0374 27.1537 31.7086 27.1537 31.7086L31.8666 29.3522C32.4573 29.0568 33.1557 29.0721 33.7329 29.393L38.1522 31.85C38.7869 32.2028 39.1804 32.8719 39.1804 33.598V38.6715C39.1804 41.2552 36.7805 43.1213 34.3325 42.2952C29.3045 40.5987 21.4998 37.3685 16.553 32.4216C11.6061 27.4748 8.37586 19.6701 6.67934 14.6422C5.85332 12.1941 7.71941 9.79422 10.3031 9.79422H15.3758Z"/><path stroke-linecap="round" d="M29 20L42 7.5"/><path stroke-linecap="round" d="M42 20H29V7"/></g>`),
		g.Group(children),
	)
}