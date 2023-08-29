package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circletwitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm301-768q-6 3-18 11l-19.5 13l-18.5 10l-21 7q-37-41-91-41q-117 0-117 98v59q-161-8-247-118q-25 26-25 57q0 66 49 100q-6 0-17 1t-17.5 0t-14.5-5q0 46 24.5 76.5T348 564q-10 12-28 12q-16 0-28-9q0 39 37.5 60.5T414 650q-18 27-52.5 40.5T288 704q-14 0-38.5-7t-25.5-7q16 32 65.5 55T415 768q67 0 125-23.5t99-62.5t70.5-89t44-103.5T768 384q0-2 12-8.5t28-17.5t24-23q-54 0-72 2q35-21 53-81z"/>`),
		g.Group(children),
	)
}