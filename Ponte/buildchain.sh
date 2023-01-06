#!/bin/bash
FUNCTION=$1
CATEGORY=$2 #osmosis_egotrade
PARAM_1=$3
PARAM_2=$4
PARAM_3=$5

Ignite() {
    git clone https://github.com/ignite/cli --depth=1
    cd cli && make install

    
    #curl https://get.ignite.com/cli! | bash
    
}

Init() {
    rm -rf pontechain
    ignite scaffold chain ponte --address-prefix pontechain

    #alice : display worry outdoor tape educate patch busy guilt lift tissue banana violin first april cup height dish spot heavy leave kite assume gauge plate
    #bob: slot grace dismiss vital organ firm library number photo amused loyal recipe ritual butter catalog win furnace believe sting gain duty ankle figure alien
    
}


#################################### End of Function ###################################################
if [[ $FUNCTION == "" ]]; then
    Ignite
else
    $FUNCTION $CATEGORY
fi
