package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envato(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.865.26c-.984-.495-3.625-.328-6.927.828c-5.609 3.958-10.229 9.568-10.724 18.813c0 .323-.656 0-.823-.167c-1.484-2.807-2.146-5.943-.823-10.563c.328-.328-.495-.823-.495-.656c-.333.161-1.484 1.479-2.313 2.802c-3.797 6.594-1.323 15.349 5.281 18.974c6.594 3.63 15.177 1.318 18.974-5.448c4.458-7.75.495-23.094-2.151-24.583z"/>`),
		g.Group(children),
	)
}