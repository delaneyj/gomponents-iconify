package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picasa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.808 1.823A9.175 9.175 0 0 0 .8 10.021c0 .633.08 1.275.221 1.918L8.97 4.7L5.808 1.823zm7.985-.197A9.117 9.117 0 0 0 9.999.8c-.934 0-1.855.156-2.749.441l6.543 5.951V1.626zm1.399.812v10.617h3.485a9.192 9.192 0 0 0 .522-3.035c.001-3.033-1.522-5.872-4.007-7.582zM1.463 13.429a9.227 9.227 0 0 0 3.368 4.184v-7.25l-2.045 1.861c-.698.634-1.28 1.166-1.323 1.205zm4.767 4.996a9.105 9.105 0 0 0 2.92.775h1.689c3.019-.281 5.727-2.068 7.199-4.744H6.23v3.969z"/>`),
		g.Group(children),
	)
}