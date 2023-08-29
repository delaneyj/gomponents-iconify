package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagFo4x30"><path fill-opacity=".7" d="M-78 32h640v480H-78z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="0" clip-path="url(#flagFo4x30)" transform="translate(78 -32)"><path fill="#fff" d="M-78 32h663.9v480H-78z"/><path fill="#003897" d="M-76 218.7h185.9V32H216v186.7h371.8v106.6H216V512H109.9V325.3h-186V218.7z"/><path fill="#d72828" d="M-76 245.3h212.4V32h53.1v213.3H588v53.4H189.5V512h-53V298.7H-76v-53.4z"/></g>`),
		g.Group(children),
	)
}