package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emojisorry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 512q0 104-40.5 199t-109 163.5t-163.5 109t-199 40.5t-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0q157 0 286 87l66-87l113 151q39 42 45.5 101.5T1001 361q23 74 23 151zm-585 11l-89-118q-1-1-3.5-4t-3.5-3.5t-3-3t-3-3t-3-2.5t-3.5-2.5t-3.5-1.5t-4-1t-4-1q-12 0-30 24l-88 117q-9 9-9 21.5t9 21.5t22 9t22-9l75-100l75 100q9 9 22 9t22-9t9-22t-9-22zm-23 245q-13 0-22.5 9.5T384 800t9.5 22.5T416 832h192q13 0 22.5-9.5T640 800t-9.5-22.5T608 768H416zm257-361l-88 117q-9 9-9 21.5t9 21.5t22 9t22-9l75-100l75 100q9 9 22 9t22-9t9-22t-9-22l-89-118q-15-21-31-22q-12 0-30 24zm259-193l-68-86l-68 86q-28 29-28 70.5t28 70.5t68 29t68-29t28-70.5t-28-70.5z"/>`),
		g.Group(children),
	)
}