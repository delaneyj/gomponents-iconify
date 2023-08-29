package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.036 10.303a5.803 5.803 0 1 0-11.606 0c0 3.205 3.03 9.708 5.803 9.708c2.992 0 5.803-6.503 5.803-9.708Zm10.959 15.013a6.589 6.589 0 1 0-3.514-12.7c-3.507.97-9.706 6.255-8.866 9.29c.906 3.275 8.873 4.38 12.38 3.41ZM26.64 39.44a5.803 5.803 0 0 0 9.236-7.027c-1.94-2.55-8.289-5.892-10.496-4.213c-2.382 1.812-.68 8.69 1.26 11.24ZM8.52 33.597a6.459 6.459 0 0 0 10.929 6.887c1.902-3.018 2.908-10.94.296-12.586c-2.819-1.775-9.324 2.681-11.225 5.699ZM11.3 16.46a4.902 4.902 0 1 0-2.45 9.494c2.622.677 8.58-.429 9.166-2.698c.632-2.448-4.095-6.119-6.717-6.795Z"/>`),
		g.Group(children),
	)
}