package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Markdownlint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m342.177 285.75l130.945-130.942L512 190.744L341.13 360.576L237.573 255.963l38.437-36.378l66.166 66.166zm-65.694-100.538l-74.147 70.177l112.973 114.127H26.11C11.806 369.516 0 357.71 0 343.18V168.593c0-14.303 11.806-26.109 26.109-26.109h290.779c14.302 0 25.517 11.806 25.29 26.109v82.312l-65.695-65.693zm-80.626-.356h-46.03l-34.523 46.03l-34.523-46.03h-46.03v138.09h46.03v-69.045l34.523 44.19l34.523-44.19v69.046h46.03V184.856z"/>`),
		g.Group(children),
	)
}