package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M192 344v-43h43v43h-43zM213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zm0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50zm0-299q35.5 0 60.5 25t25 60q0 18-10 32.5t-22 23t-22 22t-10 29.5h-43q0-23 10-39.5t22-24t22-18.5t10-25q0-17-12.5-29.5t-30-12.5t-30 12.5T171 173h-43q0-35 25-60t60.5-25z"/>`),
		g.Group(children),
	)
}