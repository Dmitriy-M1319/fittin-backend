package llm

const requestFormatString = `Есть следующие показатели результатов теста MMPI
ScaleResult(scale: '(Hs)', value: %d),
ScaleResult(scale: '(D)', value: %d),
ScaleResult(scale: '(Hy)', value: %d),
ScaleResult(scale: '(Pd)', value: %d),
ScaleResult(scale: '(Mf)', value: %d),
ScaleResult(scale: '(Pa)', value: %d),
ScaleResult(scale: '(Pt)', value: %d),
ScaleResult(scale: '(Sc)', value: %d),
ScaleResult(scale: '(Ma)', value: %d),
ScaleResult(scale: '(Si)', value: %d),
Интерпретируй эти результаты (без форматирования в Markdown, кроме списков)`
