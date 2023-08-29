package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitlabFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.663 9.987l-.028-.072l-2.719-7.094a.71.71 0 0 0-.706-.449a.711.711 0 0 0-.654.522L15.72 8.52H8.282L6.443 2.895a.711.711 0 0 0-.652-.524a.72.72 0 0 0-.707.45L2.362 9.925l-.028.07a5.057 5.057 0 0 0 1.674 5.838l.01.007l.024.019l4.147 3.104l2.05 1.553l1.247.944a.843.843 0 0 0 1.016 0l1.247-.944l2.05-1.553l4.172-3.123l.01-.008a5.055 5.055 0 0 0 1.682-5.845Z"/>`),
		g.Group(children),
	)
}