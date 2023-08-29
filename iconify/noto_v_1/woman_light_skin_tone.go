package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomanLightSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1WomanLightSkinTone0" fill="#444" d="M49.91 77.19c0 3.81-2.55 6.9-5.71 6.9s-5.73-3.09-5.73-6.9c0-3.82 2.57-6.91 5.73-6.91c3.16 0 5.71 3.09 5.71 6.91m28.17 0c0 3.81 2.56 6.9 5.72 6.9c3.15 0 5.72-3.09 5.72-6.9c0-3.82-2.57-6.91-5.72-6.91c-3.16 0-5.72 3.09-5.72 6.91"/></defs><defs><path id="notoV1WomanLightSkinTone1" fill="#312d2d" d="M76.46 42.76c-4.34 4.17-25.9 23.6-55.62 28.68V31.6S33.67 6.15 64.45 6.15s43.61 26.09 43.61 26.09v39.2s-20.17-7.74-30.3-28.43c-.25-.51-.9-.64-1.3-.25z"/><path id="notoV1WomanLightSkinTone2" fill="#513f35" d="M76.38 100.52H51.62c-1.42 0-2.03.95-1.09 2.38c1.31 2 6.19 5.85 13.47 5.85s12.16-3.85 13.47-5.85c.94-1.43.33-2.38-1.09-2.38z"/></defs><path fill="#312d2d" d="m118.5 55.8l-53.48-2.22l-1.02-.02l-1.02.02L9.5 55.8s-.27 25.78 3.16 48.84c2.55 17.15 2.99 19.87 15.19 18.47c17.52-2.01 26.24-4.48 35.66-4.5c.16 0 .33.01.49.01c.16 0 .33-.01.49-.01c9.42.01 18.15 2.49 35.66 4.5c12.2 1.4 12.64-1.32 15.19-18.47c3.43-23.06 3.16-48.84 3.16-48.84z"/><ellipse cx="64" cy="56.06" fill="#312d2d" rx="54.5" ry="50.3"/><path fill="#fadcbc" d="M64.17 14.87c-33.11 0-42.95 23.36-42.95 58.91c0 35.56 26.86 48.36 42.95 48.36c16.09 0 42.61-13.14 42.61-48.69c0-35.56-9.5-58.58-42.61-58.58z"/><use href="#notoV1WomanLightSkinTone0"/><use href="#notoV1WomanLightSkinTone1"/><use href="#notoV1WomanLightSkinTone2"/><path fill="#e59600" d="M69.98 89.99c-2.11.6-4.29.89-5.98.89c-1.69 0-3.87-.29-5.98-.89c-.9-.26-1.25.6-.93 1.17c.67 1.18 3.36 3.55 6.91 3.55c3.55 0 6.24-2.37 6.91-3.55c.33-.57-.03-1.42-.93-1.17z"/><use href="#notoV1WomanLightSkinTone1"/><use href="#notoV1WomanLightSkinTone0"/><use href="#notoV1WomanLightSkinTone2"/><path fill="#dba689" d="M69.98 89.99c-2.11.6-4.29.89-5.98.89c-1.69 0-3.87-.29-5.98-.89c-.9-.26-1.25.6-.93 1.17c.67 1.18 3.36 3.55 6.91 3.55c3.55 0 6.24-2.37 6.91-3.55c.33-.57-.03-1.42-.93-1.17z"/>`),
		g.Group(children),
	)
}