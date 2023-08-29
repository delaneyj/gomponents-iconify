package skill_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GtkDark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none"><rect width="256" height="256" fill="#242938" rx="60"/><path fill="#E40000" stroke="#fff" stroke-width="4.919" d="M30.875 169.163L26 52.061l77.923 56.207v126.477l-73.048-65.582Z"/><path fill="#7FE719" stroke="#fff" stroke-width="4.919" d="m103.946 234.767l116.896-46.831l9.75-117.101l-126.646 37.479V234.79v-.023Z"/><path fill="#729FCF" stroke="#fff" stroke-width="4.919" d="m26 52.062l77.923 56.206L230.568 70.79L143.606 21L26 52.062Z"/><path stroke="#fff" stroke-width="2.901" d="M143.227 135.325V21.546m0 113.779L31.3 167.183l111.927-31.858Zm0 0l78.089 52.338l-78.089-52.338Z"/></g>`),
		g.Group(children),
	)
}