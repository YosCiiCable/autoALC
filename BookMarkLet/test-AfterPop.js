javascript:(
	function(){
		var T=prompt('更新間隔?(秒)','60');
		if(T &&! is NaN(T)){
			var F='
				<html>
					<frameset rows="*,0">
						<frame src="'+location+'"><frame>
					</frameset>
				</html>
			';
			var W=open();
			with(W.document){
				write(F);
				close();
			}
			var H='
				<html>
					<script>
						function R(){
							parent.frames[0].location="'+location+'";
						}
						setInterval("R()",'+T*1000+');
					</script>
				</html>
			';
			with(W.frames[1].document){
				write(H);
				close();
			}
		}
	}
)();


