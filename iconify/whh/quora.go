package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quora(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M708 871q26 25 60 25q27 0 45.5 19t18.5 45.5t-18.5 45T768 1024q-43 0-96-25t-89-61q-66 22-135 22q-91 0-174-38T131 819.5t-95.5-153T0 480t35.5-186.5t95.5-153T274 38T448 0t174 38t143 102.5t95.5 153T896 480q0 119-50.5 221.5T708 871zM448 128q-70 0-128.5 47t-93 128T192 480t34.5 177t93 128T448 832q41 0 80-18q-29-46-80-46q-26 0-45-18.5t-19-45t19-45.5t45-19q43 0 95.5 24.5T632 725q72-102 72-245q0-96-34-177t-93-128t-129-47z"/>`),
		g.Group(children),
	)
}