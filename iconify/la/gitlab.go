package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitlab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m8.383 1.973l-4.305 11.48l-.242.652l-2.04 5.438L16 29.875l14.203-10.332l-2.039-5.438l-4.55-12.132l-3.731 11.48h-7.766zM8.25 8.027l1.766 5.426H6.215zm15.5 0l2.035 5.426h-3.8zM5.465 15.453h5.2l3.429 10.563l-9.89-7.196zm7.3 0h6.47L16 25.403zm8.57 0h5.196l1.266 3.367l-9.895 7.196z"/>`),
		g.Group(children),
	)
}