package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TongueNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsTongueNegative0)"><path fill-rule="evenodd" d="M40 24c0 6.84-4.292 12.677-10.33 14.966c.214-.616.33-1.277.33-1.966v-1.272c1.158-.655 2.429-2.379 3.4-3.917c.793-1.258-.15-2.811-1.636-2.811H16.236c-1.487 0-2.429 1.553-1.635 2.81c.97 1.54 2.241 3.263 3.399 3.918V37c0 .689.116 1.35.33 1.966C12.292 36.677 8 30.84 8 24c0-8.837 7.163-16 16-16s16 7.163 16 16Zm-9.5 2c1.38 0 2.5-1.79 2.5-4s-1.12-4-2.5-4s-2.5 1.79-2.5 4s1.12 4 2.5 4ZM20 22c0 2.21-1.12 4-2.5 4S15 24.21 15 22s1.12-4 2.5-4s2.5 1.79 2.5 4Z" clip-rule="evenodd"/><path d="M20 31v6a4 4 0 0 0 8 0v-6h-3v4a1 1 0 1 1-2 0v-4h-3Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0Zm-6 24c0 8.611-6.047 15.81-14.126 17.582A5.976 5.976 0 0 1 24 43a5.976 5.976 0 0 1-3.874-1.418C12.047 39.81 6 32.612 6 24c0-9.941 8.059-18 18-18s18 8.059 18 18Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsTongueNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}