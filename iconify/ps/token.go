package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Token(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256q0 100 67.5 173T235 510v2h42v-2q93-8 159.5-74T510 277h2v-42h-2q-8-100-81-167.5T256 0zm21 45q62 6 113 47l-44 45q-29-22-69-28V45zm-42 0v64q-35 3-69 28l-44-45q47-41 113-47zM92 122l45 44q-25 34-28 69H45q3-63 47-113zM45 277h64q6 33 25 64l-51 39q-35-50-38-103zm190 190q-68-6-124-55l51-39q30 24 71 30v64h2zm21-104q-45 0-76-31t-31-76t31-76t76-31t76 31t31 76t-31 76t-76 31zm21 104v-64q58-6 92-51l51 38q-56 71-143 77zm169-113l-51-38q8-26 10-39h64q-6 41-23 77zm-43-119q-6-40-28-69l45-44q41 51 47 113h-64zm-104 21q0 18-12.5 30.5T256 299t-30.5-12.5T213 256t12.5-30.5T256 213t30.5 12.5T299 256z"/>`),
		g.Group(children),
	)
}