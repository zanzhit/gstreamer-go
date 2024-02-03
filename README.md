## Данная программа может записывать, смешивать видео и аудио дорожки rtsp камер и mkv файлов.

### Для работы программы необходима программа gstreamer и подключенный VPN к необходимым камерам, если импользуете запись камер вне локальной сети.

#### При запуске программа будет ждать пользовательского ввода. Возможные варианты ввода:
1. Ввод названия txt файла. Если файл находится вне директории программы, то укажите полный путь. Файл должен содержать только rtsp-ссылки камер. В строке может находиться 1 или 2 ссылки. Если в строке находится 1 ссылка, то ведется запись видео с одной камеры наблюдения, если указаны две ссылки, то из первой ссылки берется видеодорожка, а из второй аудиодорожка.

2. Ввод названия json файла. Файл должен быть следующего вида:
{
    "arraySources": [
        {
            "videoSource": "",
            "audioSource": ""
        }
    ]
}
В videoSource указывается источник видеодорожки, в audioSource источник аудиоданных. Данные поля могут принимать rtsp-ссылки (тогда начинается смешанная запись с камер), либо mkv файлы (тогда два файла смешиваются).

3. Ввод 1 rtsp-ссылки. При вводе одного адреса камеры ведется обычная запись с камеры.
4. Ввод 2 rtsp-ссылок через пробел. При вводе формата: "rtsp://адрес rtsp://адрес", идет смешанная запись с двух камер. Из первой ссылки берется видедорожка, из второй аудиодорожка.
5. Ввод 2 mkv файлов. При указании названий двух mkv файлов, из первого возьмется видеодорожка, из второго аудиодорожка. 

**В названии json, txt, mkv файлов не должно быть пробелов.**

Для остановки записи нажмите Enter.