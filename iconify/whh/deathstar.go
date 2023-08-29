package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deathstar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1010 629q-202 75-498 75T14 629Q0 569 0 512q0-104 40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199q0 57-14 117zM443 143.5Q389 115 323.5 137t-103 81.5T195 341t66 91.5t119.5 6.5t103-81.5T509 235t-66-91.5zM352 320q-13 0-22.5-9.5T320 288t9.5-22.5T352 256t22.5 9.5T384 288t-9.5 22.5T352 320zm650 341q-49 160-184 261.5T512 1024T206 922.5T22 661q203 107 490 107t490-107z"/>`),
		g.Group(children),
	)
}