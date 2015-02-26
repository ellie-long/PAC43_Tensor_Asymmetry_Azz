#!/bin/sh
FILE=azz
VER=01


latex ${FILE}
bibtex ${FILE}
latex ${FILE}
latex ${FILE}
#dvipdf ${FILE}.dvi
dvips -t letter -Ppdf ${FILE}.dvi -o
ps2pdf ${FILE}.ps
rm ${FILE}.log ${FILE}.dvi ${FILE}.aux ${FILE}.toc ${FILE}.ps
rm ${FILE}.bbl ${FILE}.blg

mv ${FILE}.pdf ${FILE}_v${VER}.pdf

#evince ${FILE}.pdf
#open ${FILE}_v${VER}.pdf
