package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HematologyLaboratory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Zm11.8 10h7.4v1c0 .285.223.721.65 1.204a6.096 6.096 0 0 0 .744.706l.008.006l.398.3V21h-2v-.82a8.019 8.019 0 0 1-.65-.651c-.317-.36-.727-.893-.962-1.529h-3.843a4.55 4.55 0 0 1-.157.3c-.205.36-.465.716-.706 1.02c-.243.306-.485.579-.664.774l-.018.02V31.5a3.5 3.5 0 0 0 6.663 1.5h2.13A5.5 5.5 0 0 1 19 31.5V19.302l.29-.293l.003-.002l.012-.013l.051-.052l.189-.2c.158-.173.366-.41.572-.668c.21-.263.4-.529.532-.762c.103-.182.136-.282.146-.313l.004-.011v.001l.001.004V16ZM32 28.09c0 1.608-1.343 2.91-3 2.91s-3-1.302-3-2.91c0-2.544 3-5.09 3-5.09s3 2.546 3 5.09ZM29 10h-9v4h9v-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}