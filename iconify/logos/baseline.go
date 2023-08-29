package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baseline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#99DBB1" d="m182.918 35.045l38.037 38.036l-37.61 37.611l-23.505-23.506l-17.523 17.521l23.506 23.506l17.522 17.523L256 73.081L238.477 55.56l-38.036-38.037z"/><path fill="#99DBB3" d="m72.655 0l-55.56 55.559l17.523 17.523l38.037-38.037l23.506 23.506l17.522-17.522l-23.506-23.506z"/><path fill="#008F32" d="M73.082 110.691L17.523 55.132L0 72.654l73.082 73.082l17.095-17.095l.428-.427L200.868 17.95L183.346.427z"/>`),
		g.Group(children),
	)
}