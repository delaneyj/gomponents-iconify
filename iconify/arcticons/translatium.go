package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Translatium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.25 36.476c3.607 0 7.02.585 9.75 1.56V11.224c-2.828-.975-6.143-1.56-9.75-1.56s-7.02.585-9.75 1.56v26.715c2.73-.975 6.045-1.463 9.75-1.463Zm9.75 1.56c2.828-.975 6.143-1.56 9.75-1.56s7.02.585 9.75 1.56V11.224c-2.828-.975-6.143-1.56-9.75-1.56s-7.02.585-9.75 1.56"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.777 24.679h9.945l-4.68 5.655m-1.364-9.458c.39.195.877.488 1.267.878c.585.487.975.975 1.365 1.365m-7.995-.39c1.17-.683 2.535-1.56 3.998-2.828c1.17-1.072 2.047-2.047 2.827-2.925"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.477 22.534c-.974-.488-2.34-1.17-3.607-2.34c-1.268-1.073-2.243-2.243-2.828-3.12m-15.307 13.65l-4.387-13.748l-4.68 13.748m1.559-4.68h5.948"/>`),
		g.Group(children),
	)
}