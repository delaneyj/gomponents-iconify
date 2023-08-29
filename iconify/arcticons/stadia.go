package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stadia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 19.323a33.5 33.5 0 0 1 39-3.298"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.157 21.825a33.303 33.303 0 0 0-28.269-.636a33.468 33.468 0 0 1 6.775-.779a33.33 33.33 0 0 1 18.093 5.305m-23.28 5.045c4.431-2.508 8.19-3.988 13.603-4.243a33.579 33.579 0 0 0-6.982-.735a27.667 27.667 0 0 0-12.723 3.02m8.729 7.931a33.331 33.331 0 0 1 17.547-4.961q.634 0 1.262.023"/><path fill="none" stroke="currentColor" d="m4.5 19.323l3.875 9.48M43.5 16.025l-2.343 5.8m-3.384 3.907l-1.813 6.066M14.476 30.76l2.626 5.972"/>`),
		g.Group(children),
	)
}