package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<defs><path id="solarHomeBoldDuotone0" d="M10.75 9.5a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0Z"/></defs><g fill="currentColor"><path fill-rule="evenodd" d="m21.532 11.586l-.782-.626v10.29H22a.75.75 0 0 1 0 1.5H2a.75.75 0 1 1 0-1.5h1.25V10.96l-.781.626a.75.75 0 1 1-.937-1.172l8.125-6.5a3.75 3.75 0 0 1 4.686 0l8.125 6.5a.75.75 0 1 1-.936 1.172ZM12 6.75a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5Zm1.746 6.562c-.459-.062-1.032-.062-1.697-.062h-.098c-.665 0-1.238 0-1.697.062c-.491.066-.963.215-1.345.597s-.531.854-.597 1.345c-.062.459-.062 1.032-.062 1.697v4.299h7.5v-4.423c0-.612-.004-1.143-.062-1.573c-.066-.491-.215-.963-.597-1.345s-.853-.531-1.345-.597Z" clip-rule="evenodd"/><g fill-rule="evenodd" clip-rule="evenodd" opacity=".5"><use href="#solarHomeBoldDuotone0"/><use href="#solarHomeBoldDuotone0"/></g><path d="M12.05 13.25c.664 0 1.237 0 1.696.062c.492.066.963.215 1.345.597s.531.853.597 1.345c.058.43.062.96.062 1.573v4.423h-7.5v-4.3c0-.664 0-1.237.062-1.696c.066-.492.215-.963.597-1.345s.854-.531 1.345-.597c.459-.062 1.032-.062 1.697-.062h.098ZM16 3h2.5a.5.5 0 0 1 .5.5v4.14l-3.5-2.8V3.5A.5.5 0 0 1 16 3Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}