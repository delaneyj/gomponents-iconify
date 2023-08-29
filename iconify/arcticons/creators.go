package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Creators(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.925 14.164a13.98 13.98 0 0 0-9.87-4.076c-7.709 0-13.957 6.229-13.957 13.912c0 7.683 6.248 13.912 13.957 13.912a13.98 13.98 0 0 0 9.868-4.074m5.384-25.04A21.606 21.606 0 0 0 27.055 2.5c-11.913 0-21.57 9.626-21.57 21.5s9.657 21.5 21.57 21.5a21.61 21.61 0 0 0 15.252-6.296M31.54 19.528a6.355 6.355 0 0 0-4.485-1.851c-3.504 0-6.344 2.83-6.344 6.323a6.31 6.31 0 0 0 1.857 4.47"/>`),
		g.Group(children),
	)
}