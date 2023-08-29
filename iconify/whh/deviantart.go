package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deviantart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-768-64h768q27 0 45.5-18.5t18.5-45.5V613l-320 91l-160-256h160l68 101l219-58q-155-107-351-107q-94 0-183 27l183 325l-512 122v38q0 27 19 45.5t45 18.5zm238-319l-96-181q-48 24-78.5 47.5t-43.5 47.5t-16.5 41.5t-3.5 43.5q0 29 2 63zm530-577h-768q-26 0-45 18.5t-19 45.5v271q70-45 136-71l-72-136h128l61 97q128-33 259-33q203 0 384 76V128q0-27-18.5-45.5t-45.5-18.5z"/>`),
		g.Group(children),
	)
}