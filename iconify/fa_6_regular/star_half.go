package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M293.3.6c10.9 2.5 18.6 12.2 18.6 23.4v384.7c0 8.9-4.9 17-12.7 21.2L151 509.1c-8.1 4.3-17.9 3.7-25.3-1.7s-11.2-14.5-9.7-23.5l26.2-155.6L31.1 218.2c-6.5-6.4-8.7-15.9-5.9-24.5s10.3-14.9 19.3-16.3l153.2-22.6l68.6-141.3c4.9-10.1 16.1-15.4 27-12.9zm-29.4 127.8l-28.6 58.8c-3.5 7.1-10.2 12.1-18.1 13.3L99 217.9l85.9 85.1c5.5 5.5 8.1 13.3 6.8 21l-20.3 119.7l92.5-49.4V128.4z"/>`),
		g.Group(children),
	)
}