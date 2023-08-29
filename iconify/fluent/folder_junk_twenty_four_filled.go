package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderJunkTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M17.5 11a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11zm3.31 3.252l-5.558 5.557a4 4 0 0 0 5.557-5.557zM19.75 6.5a2.25 2.25 0 0 1 2.229 1.938l.016.158l.005.154v3.06A6.5 6.5 0 0 0 12.023 20H4.25a2.25 2.25 0 0 1-2.245-2.096L2 17.75v-7.251l6.207.001l.196-.009a2.25 2.25 0 0 0 1.088-.393l.156-.12L13.821 6.5h5.929zm-2.25 6a4 4 0 0 0-3.31 6.248l5.558-5.557A3.981 3.981 0 0 0 17.5 12.5zM8.207 4c.46 0 .908.141 1.284.402l.156.12l2.103 1.751l-3.063 2.553l-.085.061a.75.75 0 0 1-.29.106L8.206 9L2 8.999V6.25a2.25 2.25 0 0 1 2.096-2.245L4.25 4h3.957z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}