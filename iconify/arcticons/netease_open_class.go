package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeteaseOpenClass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.767 30.713c-1.334-.794-1.313-19.47.022-20.3s16.67 8.3 16.668 9.925S20.1 31.505 18.767 30.712ZM21.73 11.28l-3.804 3.978m5.991-2.89l-6.124 6.443m8.169-5.323l-8.18 8.377M28.18 14.79L17.899 25.585m12.251-9.57l-10.93 11.48m12.9-10.167L21.54 28.45m12.544-9.66l-8.086 8.622"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.066 27.819c2.356.107 3.195.932 3.666 1.937M14.813 7.822l.064 30.661M9.235 32.97l18.36.223m9.941-1.471a1.106 1.106 0 0 1 .842-1.305a1.113 1.113 0 0 1 1.139 1.718c-.376.454-1.55 1.782-1.55 1.782l2.155-.423m-2.828 3.092a1.184 1.184 0 0 1-.903.8h0a1.227 1.227 0 0 1-1.436-.965l-.153-.78a1.227 1.227 0 0 1 .965-1.436h0a1.184 1.184 0 0 1 1.139.4m-8.488 2.348a1.224 1.224 0 0 1 2.402-.47l.376 1.92m-3.013-2.65l.612 3.121"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.82 36.483a1.224 1.224 0 1 1 2.401-.471l.377 1.92m-9.284.057l2.445-.112m-2.503-1.427l2.444-.113m-5.512 3.854l2.441-.166m-2.774-4.716l2.441-.167m-2.275 2.608l1.587-.108m-1.753-2.333l.333 4.882"/>`),
		g.Group(children),
	)
}