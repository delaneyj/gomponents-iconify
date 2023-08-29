package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dribbble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M308 358q56-40 69-107q-35-8-66-8q-17 0-34 3q19 57 31 112zm-95 29q31 0 59-11q-12-63-32-121q-49 16-87 52q-23 22-39 47q44 33 99 33zM47 221q0 60 39 106q19-28 46-53q42-38 94-55q-4-10-10-22q-67 21-151 22q-13 0-18-1v3zm93-150q-33 16-56 45t-32 64q3 1 13 1h3q70 0 131-19q-29-54-59-91zm73-17q-16 0-35 4q32 42 57 91q53-23 82-58q-45-37-104-37zm131 64q-36 41-92 66q4 8 11 25q24-4 48-4q33 0 69 8q-3-53-36-95zM213.5 7q88.5 0 151 62.5t62.5 151t-62.5 151t-151 62.5t-151-62.5T0 220.5t62.5-151T213.5 7z"/>`),
		g.Group(children),
	)
}