package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Supportalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm-128-83q63 19 128 19t128-19V733q-59 35-127.5 35T384 733v208zm-93-557H83q-19 63-19 128t19 128h208q-35-60-35-128t35-128zM640 83q-63-19-128-19T384 83v208q60-35 128-35t128 35V83zM512 320q-80 0-136 56t-56 136t56 136t136 56t136-56t56-136t-56-136t-136-56zm221 64q35 60 35 128t-35 128h208q19-63 19-128t-19-128H733z"/>`),
		g.Group(children),
	)
}