package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlevine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm256-512q-25 0-49.5-4t-51-15t-46-28.5T589 418t-13-66q0-29 21-62.5t43-33.5q23 0 43.5 19.5T704 320q0 86-14 128h63q15-46 15-136q0-51-34.5-85.5T640 192q-45 0-86.5 41T512 332q0 50 14 94t30.5 69.5T595 542t31 26t14 8q0 17-28 56.5T552.5 704T512 736q-4 0-24-18t-48.5-55.5t-55-85T339 461t-19-141h-64q0 101 24.5 190.5t59.5 146t73 98t64.5 59.5t34.5 18q7 0 36.5-22.5T613 753t63-83.5t28-93.5q41 3 64 0v-64z"/>`),
		g.Group(children),
	)
}