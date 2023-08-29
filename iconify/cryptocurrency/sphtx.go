package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sphtx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zM6 12.391h8.417v-.356H6zm0-1.035h8.417V11H6zm3.55 10.64h.338v-8.893H9.55zm.98 0h.337v-8.893h-.337zm15.232-.728l.238-.252l-3.798-4.012l-.238.251zM16.519 11l-.238.252l3.798 4.012l.238-.251zm8.55 11l.238-.252l-3.798-4.013l-.238.252zm-9.242-10.268l-.239.252l3.798 4.012l.239-.252zm5.421 4.768l-.454.48l-.238.252l-4.275 4.516l.238.252l4.275-4.516l.239-.252l.454-.48l.238-.252L26 11.984l-.238-.252l-4.275 4.516zm-.454-.984l-.238.252l-.455.48l-.238.252l-4.275 4.516l.239.252l4.274-4.516l.239-.252l.454-.48l.238-.252l4.275-4.516L25.07 11z"/>`),
		g.Group(children),
	)
}