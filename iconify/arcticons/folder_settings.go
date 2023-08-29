package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.207 27.175c.033-.229.05-.46.052-.69a5.04 5.04 0 0 0-.052-.68l1.484-1.133a.35.35 0 0 0 .082-.443l-1.401-2.431a.34.34 0 0 0-.422-.155l-1.752.7a5.932 5.932 0 0 0-1.184-.69l-.268-1.854a.35.35 0 0 0-.35-.298h-2.782a.36.36 0 0 0-.37.278l-.258 1.854a5.653 5.653 0 0 0-1.184.69l-1.752-.7a.33.33 0 0 0-.422.154l-1.411 2.431a.35.35 0 0 0 .092.443l1.473 1.144a3.86 3.86 0 0 0 0 .68v.69l-1.473 1.153a.35.35 0 0 0-.092.454l1.432 2.441a.32.32 0 0 0 .422.144l1.751-.7c.363.28.761.513 1.185.69l.257 1.864a.35.35 0 0 0 .35.289h2.74a.36.36 0 0 0 .35-.288l.31-1.865a5.24 5.24 0 0 0 1.195-.69l1.74.7c.16.061.342 0 .433-.144l1.401-2.431a.36.36 0 0 0-.082-.453l-1.494-1.154Zm-5.212.865a1.545 1.545 0 1 1 0-3.09h0c.848 0 1.535.687 1.535 1.535h0a1.545 1.545 0 0 1-1.535 1.555h0Z"/>`),
		g.Group(children),
	)
}