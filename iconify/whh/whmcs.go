package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whmcs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M874 384H733q-34-58-92.5-93T512 256q-106 0-181 75t-75 181t75 181t181 75q70 0 128.5-35t92.5-93h141l-1 1l101 90q-16 33-34 61l-132-35q-22 26-49 49l39 131q-25 17-62 35l-93-99q-33 11-67 17l-35 133q-25 1-29 1q-16 0-42-2l-27-132q-30-6-60-17l-90 101q-34-16-61-34l35-132q-26-22-49-49L87 798q-17-26-35-62l99-93q-11-33-17-67L1 541q-1-25-1-29q0-16 2-42l132-27q6-30 17-60L50 293q16-34 34-61l132 35q22-26 49-49L226 87q26-17 62-35l93 99q33-11 67-17L483 1q26-1 29-1q16 0 42 2l27 132q30 6 60 17l90-101q34 16 61 34l-35 132q26 22 49 49l131-39q17 25 35 62l-99 93v2l1 1z"/>`),
		g.Group(children),
	)
}