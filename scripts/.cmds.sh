 rm -rf ~/.autonomy
 autonomy init autonomy --chain-id autonomy
 autonomy add-genesis-account autonomy1748gkseu9yyrlgfucndkemwtptttshwgvgpr07 10000000000000uaut
 autonomy add-genesis-account autonomy1089z23a6pkkl5dulktklcjxe8mgprgpd3cuuu5  10000000000000uaut
 autonomy gentx a-1 100000000uaut --chain-id autonomy 
 autonomy collect-gentxs
 
 
 # Community Creation
 autonomy tx nft  create-community Thub --description "test desc" --preview_uri "test image" --from a-1 --chain-id autonomy 
 
  autonomy tx nft  create-community "Sai Community" --description "test sai desc" --preview_uri "https://ipfs.io/ipfs/test-image-url" --from main --chain-id autonomy 
  
   autonomy tx nft  create-community Prashanth --description "test prashanth desc" --preview_uri "https://ipfs.io/ipfs/test-image-url" --from main --chain-id autonomy 
 

# Query all communities
http://localhost:1317/autonomy/nft/v1beta1/communities

# Query community by communityID
http://localhost:1317/autonomy/nft/v1beta1/communities/{community_id}

# Create collection
 autonomy tx nft create Krish  community12a43d89fef54b23a9ed3a55faaa5271  ""  --symbol sym --description "test desc" --from a-1 --chain-id autonomy  --preview_uri "https://ipfs.io/ipfs/test-denom-image"
 
 # Query Collections
 http://localhost:1317/autonomy/nft/v1beta1/denoms
 # Query Collection by denom id
 http://localhost:1317/autonomy/nft/v1beta1/denoms/{collection_id}
 
 #Query Collection and NFT with collection Id
 http://localhost:1317/autonomy/nft/v1beta1/collection/{collection_id}
 
 # Sell NFT 
autonomy tx nft sell {nft_denom_id} {nft_id} 10uaut --from main  --chain-id autonomy 

 # Query MarketPlace
 http://localhost:1317/autonomy/nft/v1beta1/market_place
 
 # Query each NFT, by collection, nft
 http://localhost:1317/autonomy/nft/v1beta1/market_place/denoms/{denom_id}/nfts/{nft_id}
 
 
 
