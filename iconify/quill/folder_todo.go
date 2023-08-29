package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderTodo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12.707 16.293a1 1 0 0 0-1.414 1.414l1.414-1.414ZM15 20l-.707.707a1 1 0 0 0 1.414 0L15 20Zm6.207-4.793a1 1 0 0 0-1.414-1.414l1.414 1.414Zm-9.914 2.5l3 3l1.414-1.414l-3-3l-1.414 1.414Zm4.414 3l5.5-5.5l-1.414-1.414l-5.5 5.5l1.414 1.414ZM29 24V11h-2v13h2ZM26.003 8H17v2h9.003V8ZM17 8c-.208 0-.36-.063-.552-.228c-.245-.21-.442-.477-.805-.912c-.327-.393-.755-.872-1.35-1.24C13.678 5.236 12.933 5 12 5v2c.568 0 .947.138 1.238.318c.311.194.571.465.869.822c.262.315.627.797 1.04 1.15c.463.398 1.06.71 1.853.71V8Zm-5-3H6v2h6V5ZM3 8v16h2V8H3Zm3 19h20v-2H6v2ZM6 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V5Zm23 6c0-1.655-1.338-3-2.997-3v2c.55 0 .997.446.997 1h2ZM3 24a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1H3Zm24 0a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2Z"/>`),
		g.Group(children),
	)
}