const bitVector=function(length){
	this.v=new Uint8Array(length>>3);
	this.toggleBit=(pos)=>{this.v[pos>>3]^=1<<pos%8;};
	this.getBit=(pos)=>{return this.v[pos>>3]&1<<pos%8?true:false;};
};


const bv=new bitVector(16);

for(let i=0;i<16;i++){
	bv.toggleBit(i);
	console.log(bv.getBit(i));
	bv.toggleBit(i);
	console.log(bv.getBit(i));
}
