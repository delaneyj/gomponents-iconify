package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dblp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.075.002c-.096.013-.154.092-.094.31L4.97 7.73L3.1 8.6s-.56.26-.4.85l2.45 9.159s.16.59.72.33l6.169-2.869l1.3-.61s.52-.24.42-.79l-.01-.06l-1.13-4.22l-.658-2.45l-.672-2.49v-.04s-.16-.59-.84-1L3.5.141s-.265-.16-.425-.139zM18.324 5.03a.724.724 0 0 0-.193.06l-5.602 2.6l.862 3.2l1.09 4.08l.01.06c.05.47-.411.79-.411.79l-1.88.87l.5 1.89l.04.1c.07.17.28.6.81.91l6.95 4.269s.68.41.52-.17l-1.981-7.4l1.861-.86s.56-.26.4-.85L18.85 5.42s-.116-.452-.526-.39z"/>`),
		g.Group(children),
	)
}