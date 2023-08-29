package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M680 289V82c0-37-31-68-68-68H68C31 14 0 45 0 82v207h165c26-71 95-122 175-122s149 51 175 122h165zm0 60H527v5c0 103-84 187-187 187s-187-84-187-187v-5H0v277c0 37 31 68 68 68h544c37 0 68-31 68-68V349zM59 74h34v156H59V74zm68 0h34v156h-34V74zm68 0h34v98c-11 8-23 17-34 26V74zm145 408c71 0 127-57 127-128s-56-127-127-127s-128 56-128 127s57 128 128 128zM459 74h161v156H512c-14-21-33-38-53-52V74z"/>`),
		g.Group(children),
	)
}