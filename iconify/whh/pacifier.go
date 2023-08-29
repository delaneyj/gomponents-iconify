package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pacifier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024.695 736q0 119-84.5 203.5t-203.5 84.5q-71 0-132-32t-101-88q-95 47-172 55t-111-27q-40-39-23.5-134.5t82.5-205.5q62 48 137 48q93 0 158.5-65.5t65.5-158.5q0-75-48-137q110-66 205.5-82.5t134.5 23.5q35 34 27 111t-55 172q56 40 88 101t32 132zm-187-123q-46 66-102 122t-122 102q49 59 123 59q66 0 113-47t47-113q0-74-59-123zm-421-37q-30 0-53.5-10t-36-22t-31-22t-39.5-10q-106 0-181-75t-75-181t75-181t181-75t181 75t75 181q0 21 10 39.5t22 31t22 36t10 53.5q0 66-47 113t-113 47z"/>`),
		g.Group(children),
	)
}